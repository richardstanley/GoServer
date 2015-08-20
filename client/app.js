window.host = 'localhost';

var app = window.energyDataApp = angular.module('energyDataApp', [])

app.config(function ($stateProvider, $urlRouterProvider){
	$urlRouterProvider.otherwise('/');
})