var app = angular.module("search", []);

app.controller("SuggestCtrl", function($scope, $http) {
    $scope.loading = true;
    // hide the fact that its loading initially
    
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
            $scope.loading = false;       
            var max = $scope.max;
            // if the max somehow ranges further than 10-50, fix it
            if (max > 50) {
                max = 50;
                $scope.max = 50
            } else if (max < 10) {
                max = 10;
                $scope.max = 10;
            }
            $http({
                method: "get",
                url: "/api/v1/suggestion?q=" + $scope.search + "&max=" + max,
                timeout: 5000
            }).success(
                function(response){
                    if (response.Results == null) {
                        // there are no results found
                        $scope.noResult = true;                   // hide the previous results
                        $scope.results = null;
                    } else if(response.Results != "undefined") {
                        // there are results
                        $scope.noResult = false;
                        $scope.results = response.Results;
                        // set to the correct linking service
                        $scope.service = 
                        response.Service;
                       
                        $scope.loading = false;
                        
                    } else {
                        // tell the user that it's loading
                        $scope.loading = true;
                        console.log(response, "Results are still loading, stay tuned")
                    }
                    
                }).error(function(response) {
                    $scope.isLoading = true;
                    console.log(response, "Results are still loading, stay tuned")
                });
                
    }
});
   