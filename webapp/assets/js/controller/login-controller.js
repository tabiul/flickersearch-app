'use strict'
flickerSearchApp.controller('LoginController', function ($scope, $http, $location, AuthService) {
	$scope.login = function(){
	   var data = {
		   username : $scope.username,
		   password : $scope.password
	   }	
	   var res = $http.post('http://localhost:8080/authenticate', angular.toJson(data));
	   res.success(function(data, status, headers, config) {
			if(status == 200){
				AuthService.setUser($scope.username);
				$location.path("/images")
			}
		});
		res.error(function(response, status, headers, config) {
			$scope.message = response
		});
    }

});