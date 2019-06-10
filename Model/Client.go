package model

import (
	"log"
	"net"

	apiService "../Services/API"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (obj *Model) OpenLock(badgeID string) bool {
	var b badge
	result := true

	obj.db.Where("id = ?", badgeID).First(&b)

	if b.ID != badgeID {
		result = false
	}

	tdb := obj.db.Begin()
	defer tdb.Rollback()
	l := logs{Code: badgeID, Success: result, Mode: "READ"}
	if err := tdb.Create(&l).Error; err != nil {
		result = false
	}
	tdb.Commit()

	return result
}

func (obj *Model) AddBadge(badgeID string, nom string, prenom string) bool {
	b := badge{ID: badgeID, FirstName: nom, LastName: prenom}
	result := true

	tdb := obj.db.Begin()
	defer tdb.Rollback()

	if err := tdb.Create(&b).Error; err != nil {
		result = false
	}

	l := logs{Code: badgeID, Success: result, Mode: "ADD"}
	if err := tdb.Create(&l).Error; err != nil {
		result = false
	}
	tdb.Commit()

	return result
}

func (obj *Model) DeleteBadge(badgeID string) bool {
	obj.db.Unscoped().Where("id LIKE ?", badgeID).Delete(badge{})
	l := logs{Code: badgeID, Success: true, Mode: "DELETE"}
	obj.db.Create(&l)
	return true
}

func (obj *Model) GetServerAddress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func (obj *Model) ChangeMode(mode string) bool {
	serverAddr := obj.GetServerAddress()
	s := settings{Server: serverAddr.String(), Mode: mode}

	tdb := obj.db.Begin()
	defer tdb.Rollback()

	if err := tdb.Where(settings{Server: serverAddr.String()}).Assign(settings{Mode: mode}).FirstOrCreate(&s).Error; err != nil {
		return false
	}
	tdb.Commit()

	return true
}

func (obj *Model) GetCurrentMode() string {
	serverAddr := obj.GetServerAddress()
	var s settings
	obj.db.Where("server = ?", serverAddr.String()).First(&s)
	if s.Server == serverAddr.String() {
		return s.Mode
	}
	return ""
}

func (obj *Model) GetLastLog() string {
	var log logs
	obj.db.Last(&log)
	var bdg badge
	obj.db.Where("ID = ?", log.Code).First(&bdg)
	var res string

	switch log.Mode {
	case "ADD":
		res = "Ajout de " + bdg.FirstName + " " + bdg.LastName + " (" + log.Code + ")"
	case "DELETE":
		res = "Suppression de " + bdg.FirstName + " " + bdg.LastName + " (" + log.Code + ")"
	default:
		res = bdg.FirstName + " " + bdg.LastName + " (" + log.Code + ")" + " a badgé, déverrouillage de la porte"
	}

	switch log.Success {
	case true:
		res += " avec succès !"
	default:
		res += " échoué !"
	}

	return res
}

func (obj *Model) GetBadgesList() (dst []apiService.Badge) {
	var bdgs []badge
	obj.db.Order("created_at desc", false).Find(&bdgs)
	logObj(bdgs)

	for i := range bdgs {
		var vDTO apiService.Badge
		bdgs[i].toDTO(&vDTO)
		dst = append(dst, vDTO)
	}

	return
}

func (obj *Model) SetNomPrenom(badgeID string, nom string, prenom string) bool {

	obj.db.Model(badge{}).Where("ID = ?", badgeID).Update("last_name", prenom).Update("first_name", nom)

	return true
}
