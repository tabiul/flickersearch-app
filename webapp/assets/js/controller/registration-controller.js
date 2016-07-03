'use strict'
flickerSearchApp.controller('RegistrationController', function ($scope, $http, $location, $rootScope) {
	$scope.register = function(){
	   var data = {
		   username : $scope.username,
		   password : $scope.password
	   }	
	   var res = $http.post('http://localhost:8080/user', angular.toJson(data));
	   res.success(function(response, status, headers, config) {
			if(status == 201){
				$rootScope.username = $scope.username
				$location.path("/images")
			}
		});
		res.error(function(response, status, headers, config) {
			$scope.message = response
		});
    }
});