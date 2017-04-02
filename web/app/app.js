'use strict';

/**
 * @ngdoc overview
 * @name catapi
 * @description
 * # catapi
 *
 * Main module of the catapi application.
 */
angular
  .module('catapi', [
    'ngRoute'
  ])
  .config(function ($routeProvider) {
    console.log('routing');
    $routeProvider
      .when('/static/', {
        templateUrl: 'index.html',
        controller: 'IndexCtrl'
      })
  });
