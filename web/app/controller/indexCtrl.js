'use strict';

/**
 * @ngdoc function
 * @name catapi.controller:IndexCtrl
 * @description
 * # IndexCtrl
 * Controller of the index page of catapi
 */

angular.module('catapi',[])
  .controller('IndexCtrl', function($scope,$http) {
        $scope.catImage = "";
        console.log("controller is loaded");

        $scope.getCat = function(){
          console.log("get incoming....");
          $http.get('http://localhost:8080/cats')
          .then(function(response) {
              console.log(response.data);
              $scope.catImage = response.data.pic;
              console.log($scope.catImage);
          });
        };

  });
