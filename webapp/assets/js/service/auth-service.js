'use strict'
flickerSearchApp.factory('AuthService', function(){
var user = "";
return{
    setUser : function(aUser){
        user = aUser;
    },
    isLoggedIn : function(){
        return user !== ""
    },
	getUser: function(){
		return user
	}
  }
})