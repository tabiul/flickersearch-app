'use strict'
var flickerSearchApp = angular.module('flickrSearchApp', ['ngRoute']);

flickerSearchApp.config(function($routeProvider) {
  $routeProvider
   .when('/', {
    templateUrl: 'assets/template/login.template.html',
    controller: 'LoginController',
   })
  .when('/registration', {
    templateUrl: 'assets/template/registration.template.html',
    controller: 'RegistrationController'
  })
  .when('/images', {
    templateUrl: 'assets/template/image.template.html',
    controller: 'ImageController'
  });
});
