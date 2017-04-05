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
        $scope.catId = "";

        $scope.getCat = function(){
          var url = 'http://localhost:8080/cats';
          if($scope.catId !== "") {
            url = url + '/' + $scope.catId;
          }
          $http.get(url)
          .then(function(response) {
              $scope.catImage = response.data.pic;
          });
        };

  });
