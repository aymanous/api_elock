package apiService

import (
	helperhttp "../../Helper/Http"
)

func routes(model dao) helperhttp.Routes {
	handler := newHandler(model)

	return helperhttp.Routes{
		helperhttp.Route{
			Name:    "GetBadgeStatus",
			Method:  "GET",
			Pattern: "/status/{id}",
			Handler: helperhttp.ErrorFnHandler(handler.openLock),
		},
		helperhttp.Route{
			Name:    "AddNewBadge",
			Method:  "GET",
			Pattern: "/add/{id}",
			Handler: helperhttp.ErrorFnHandler(handler.addBadge),
		},
		helperhttp.Route{
			Name:    "DeleteBadge",
			Method:  "GET",
			Pattern: "/delete/{id}",
			Handler: helperhttp.ErrorFnHandler(handler.deleteBadge),
		},
		helperhttp.Route{
			Name:    "GetServerAddress",
			Method:  "GET",
			Pattern: "/server/address",
			Handler: helperhttp.ErrorFnHandler(handler.getServerAddress),
		},
		helperhttp.Route{
			Name:    "ChangeMode",
			Method:  "GET",
			Pattern: "/mode/{mode}",
			Handler: helperhttp.ErrorFnHandler(handler.changeMode),
		},
		helperhttp.Route{
			Name:    "GetCurrentMode",
			Method:  "GET",
			Pattern: "/mode",
			Handler: helperhttp.ErrorFnHandler(handler.getCurrentMode),
		},
		helperhttp.Route{
			Name:    "GetLastLog",
			Method:  "GET",
			Pattern: "/log/last",
			Handler: helperhttp.ErrorFnHandler(handler.getLastLog),
		},
		helperhttp.Route{
			Name:    "GetBadgesList",
			Method:  "GET",
			Pattern: "/badges",
			Handler: helperhttp.ErrorFnHandler(handler.getBadgesList),
		},
		helperhttp.Route{
			Name:    "SetNomPrenom",
			Method:  "GET",
			Pattern: "/badges/{code}",
			Handler: helperhttp.ErrorFnHandler(handler.setNomPrenom),
		},
	}
}
