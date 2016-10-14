var React = require('react');
var ReactRouter = require('react-router');
var Router = ReactRouter.Router;
var Route = ReactRouter.Route;
var hashHistory = ReactRouter.hashHistory;

var MainContainer = require('../containers/MainContainer');
var MovieContainer = require('../containers/MovieContainer');


var someAuthCheck = function(nextState, transition) { console.log('here')}

var routes = (
  <Router history={hashHistory}>
    <Route path='/' component={MainContainer} onEnter={someAuthCheck}>
      <Route path='/movies' component={MovieContainer} />
    </Route>
  </Router>
)

module.exports = routes
