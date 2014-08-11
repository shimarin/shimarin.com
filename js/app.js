angular.module("MyApp", [])
.run(["$rootScope", function($scope) {
    $scope.message = "Hello! World!";
}]);
