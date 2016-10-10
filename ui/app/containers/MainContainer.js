var React = require('react');

var Navigation = require('./NavigationContainer');

var MainContainer = React.createClass({
  render: function() {
    return (
      <div id='wrapper'>
        <Navigation />
        {this.props.children}
      </div>
    )
  }
});

module.exports = MainContainer;
