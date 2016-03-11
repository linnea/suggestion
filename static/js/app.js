var app = angular.module("search", []);

app.controller("SuggestCtrl", function($scope, $http) {
    $scope.isLoading = false;
    $scope.max = 20;
    $scope.$watch('search', function() {
        fetch($scope);
    });
    $scope.$watch('max', function() {
        fetch($scope);
    });
    $scope.submit = function() {
        fetch($scope);
    }
    function fetch($scope) {
            var max = $scope.max;
            if (max > 50) {
                max = 50;
                $scope.max = 50
            } else if (max < 10) {
                max = 10;
                $scope.max = 10;
            }
            var uri = "/api/v1/suggestion?q=" + $scope.search + "&max=" + max;
            $http({
                method: "get",
                url: uri,
                timeout: 5000
            }).success(
                function(response){
                    $scope.isLoading = false;
                    $scope.results = response.Results;
                    $scope.service = 
                    response.Service;
                }).error(function(response) {
                    $scope.isLoading = true;
                    console.log(response, "Results are still loading, stay tuned")
                });
                
    }
    
    /*
        $http.post('/api/v1/suggestion/', { "data" : $scope.keywords}).
		success(function(data, status) {
			$scope.status = status;
			$scope.data = data;
			$scope.result = data; // Show result from server in our <pre></pre> element
		})
		.
		error(function(data, status) {
			$scope.data = data || "Request failed";
			$scope.status = status;			
		});*/
});
    
    /*
    $http({method: 'GET', url: '/api/v1/suggestion/' + $routeParams.q + $routeParams.max}).success(function(data) {
      // RELAY THAT INFO
      controller.results = data;
      
  });*/