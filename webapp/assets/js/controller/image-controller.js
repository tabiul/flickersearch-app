'use strict'
flickerSearchApp.controller('ImageController', function ($scope, $http, AuthService, $location) {
  
	function processJSON(data){
	   $scope.pages = data.pages
	   $scope.page = data.page
       $scope.imageURLList = data.images;
	}
	
	$scope.init = function(){
		if(!AuthService.isLoggedIn()){
			$location.path("/");
			return
		}
		$scope.username = AuthService.getUser()
		$http.get('http://localhost:8080/history?username=' + $scope.username).success(function(data) {
			if(data !== ""){
              $scope.history = angular.fromJson(data);		
			}
       });
	}
    $scope.search = function(){
	   var res = $http.get('http://localhost:8080/image?search=' + $scope.query + "&username=" + $scope.username);
	   
	   res.success(function(data) {
         processJSON(angular.fromJson(data));
         if($scope.history !== undefined ){
		   $scope.history.push(search)
	     }		 
       });
	   
	   res.error(function(response, status, headers, config){
		  $scope.message = response 
	   });
	   
	   
    }

    $scope.next = function(){
	   if($scope.page < $scope.pages){
		 $http.get('http://localhost:8080/image?search=' + $scope.query + "&page=" + ($scope.page + 1)).success(function(data) {
         processJSON(angular.fromJson(data));	
       });  
	  }	
	}

    $scope.prev = function(){
	   if($scope.page > 1){
		 $http.get('http://localhost:8080/image?search=' + $scope.query + "&page=" + ($scope.page - 1)).success(function(data) {
         processJSON(angular.fromJson(data));	
       });  
	  }		
	}

    $scope.complete = function(){
		if($scope.history !== undefined && $scope.history.length > 0){
          $("#search").autocomplete({
            source: $scope.history
          });
		}
    }

    $scope.logout = function(){
		AuthService.setUser("");
		$location.path("/");
	}	
  });