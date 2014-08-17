<div class="modal-header">
  <button type="button" class="close" ng-click="$dismiss()" aria-hidden="true">×</button>
  <h3>{{wikipedia.title}} <small>Wikipediaより</small></h3>
</div>
<div class="modal-body">
  {{wikipedia.summary}}
</div>
<div class="modal-footer">
  <a class="btn btn-default" href="{{wikipedia.url}}" ng-click="$close()" target="_blank">Wikipediaで詳しく見る</a>
  <button class="btn btn-primary" ng-click="$close()">閉じる</button>
</div>
