angular.module("MyApp", ["ui.bootstrap", "angulartics", "angulartics.google.analytics"])
.controller("NetKjkController", ["$scope", "$modal", "$analytics", function($scope, $modal, $analytics) {
    $scope.about = function() {
	$analytics.eventTrack("click", {category:"button", label:"AboutNetKjk"});
        $modal.open({templateUrl:"about.html"});
    }
}])
.controller("OptoutController", ["$scope", function($scope) {
    var _gaq = _gaq || []
    _gaq.push(['_setVar', 'no_analytics']);
    // TBD
    // https://productforums.google.com/forum/#!topic/analytics/ftLKh-fsUws
}])
.run(["$rootScope", "$analytics", "$location", function($scope, $analytics, $location) {

    $analytics.pageTrack($location.path());
}]);
