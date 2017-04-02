'use strict';

/**
 * @ngdoc function
 * @name catapi.controller:IndexCtrl
 * @description
 * # IndexCtrl
 * Controller of the index page of catapi
 */
angular.module('catapi')
  .controller('IndexCtrl', function ($scope) {
        $scope.catImage = "";

        $scope.getCat = function(){
          console.log("get incoming....");
          $http.get('http://localhost:8080/cats')
          .then(function(response) {
              $scope.catImage = response.data.pic;
          });
        };
  });
