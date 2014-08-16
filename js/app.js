angular.module("MyApp", ["ui.bootstrap", "angulartics", "angulartics.google.analytics"])
.controller("NetKjkController", ["$scope", "$modal", function($scope, $modal) {
    $scope.about = function() {
        $modal.open({templateUrl:"about.html"});
    }
}])
.controller("OptoutController", ["$scope", function($scope) {
    var _gaq = _gaq || [];
    console.log(_gaq);
    _gaq.push(['_setVar', 'no_analytics']);
}])
.run(["$rootScope", function($scope) {
    $scope.message = "Hello! World!";
}]);
