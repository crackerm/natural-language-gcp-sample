
function AnalyzeCtrl($scope, $http) {
  $scope.result = '';
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/testnatural/').
      success(function(data) { $scope.result = data; }).
      error(logError);
  };

  $scope.addText = function() {
    $scope.working = true;

    $http.post('/testnatural/', {Analyzetext: $scope.newText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.newText = '';
        })
      });
  };

  refresh().then(function() { $scope.working = false; });
}