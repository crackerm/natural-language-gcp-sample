<!doctype html>
<html ng-app>

<head>
  <title></title>
  <script src='/static/js/angular.min.js'></script>
  <meta name='viewport' content='width=device-width, initial-scale=1.0'>
  <link href='http://fonts.googleapis.com/css?family=Roboto:400,300' rel='stylesheet' type='text/css'>
  <script src='/static/js/analyze.js'></script>
  <link rel='stylesheet' href='/static/css/analyze.css'>
</head>

<body>
<div class='container' ng-controller='AnalyzeCtrl'>
  <h2 class='charcoal rounded-box'>Google Natural Language</h1>

  <ul class='grey rounded-box'>
    <li>
      {{result.Result.Text}} - {{result.Result.Emoji}} <br/>
      Magnitude : {{result.Result.Magnitude}} <br/>
      Score : {{result.Result.Score}}
    </li>
  </ul>

  <form>
    <input type='text' class='rounded-box' placeholder='テキスト入力してください' ng-model='newText'>
    <button class='grey rounded-box' ng-click='addText()' ng-disabled='working'>確認</button>
  </form>

  <img class='spinner' src='/static/img/spinner.gif' alt='Loading' ng-class='{working: working}'/>
</div>
</body>
</html>