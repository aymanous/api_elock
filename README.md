# API ELOCK

## Etapes d'utilisation :

  1. Copie du repository Github sur votre machine :
  </br>git clone https://github.com/aymanous/api_elock.git

  2. Paramètrage de la BDD
  </br> a. rm -fr sondes.json (fichier généré à chaque compilation, et contenant les identifiants de BDD de l'ancienne compilation)
  </br> b. Dans 'Configuration/Data.go' : modifier les identifiants de BDD au niveau de 'DBConfig'

  2. Compilation du projet (génère un fichier executable nommé 'API')
  </br>go build
 
  3. Exécution du projet
  </br>./API -config eLock
  
A ce niveau, l'API écoute sur le port 8080.
Exemple : http://localhost:8080/badges (fournit la liste des badges disponibles)

## Technologies utilisées :
  - GOLANG
  - GORM (ORM GO<>MYSQL)
  - MYSQL

## Dépendances :
  Ce module du projet eLock nécessite au préalable :
  - l'installation de Go sur votre machine
  - le lancement de la base de données MySql
