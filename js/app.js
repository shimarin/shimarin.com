angular.module("MyApp", ["ngResource", "ui.bootstrap", "angulartics", "angulartics.google.analytics"])
.controller("NetKjkController", ["$scope", "$resource", "$modal", "$analytics", function($scope, $resource, $modal, $analytics) {
    $scope.about = function() {
	$analytics.eventTrack("click", {category:"button", label:"AboutNetKjk"});
        $modal.open({templateUrl:"about.html"});
    }
    $scope.wikipedia = function(name) {
	$analytics.eventTrack("click", {category:"help", label:name});
        var newScope = $scope.$new();
        newScope.wikipedia = $resource("/wikipedia/:lang/:name").get({lang:"ja",name:name})
        $modal.open({templateUrl:"wikipedia.tpl",scope:newScope});
    }
}])
.controller("OptoutController", ["$scope", function($scope) {
    window._gaq = window._gaq || []
    window._gaq.push(['_setVar', 'no_analytics']);
    // TBD
    // https://productforums.google.com/forum/#!topic/analytics/ftLKh-fsUws
}])
.run(["$rootScope", "$analytics", "$location", function($scope, $analytics, $location) {

    $analytics.pageTrack($location.path());
}]);
