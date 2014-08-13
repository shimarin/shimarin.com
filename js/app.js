angular.module("MyApp", ["ui.bootstrap"])
.controller("NetKjkController", ["$scope", "$modal", function($scope, $modal) {
    $scope.about = function() {
        $modal.open({templateUrl:"about.html"});
    }
}])
.run(["$rootScope", function($scope) {
    $scope.message = "Hello! World!";
}]);
