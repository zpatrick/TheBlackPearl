var React = require('react');
var Bootstrap = require('react-bootstrap');
var Nav = Bootstrap.Nav;
var NavItem = Bootstrap.NavItem;
var Navbar = Bootstrap.Navbar;

var NavigationContainer = React.createClass({
  render: function() {
    return (
      <Navbar>
      <Navbar.Header>
        <Navbar.Brand>
          <a href='#/'>The Black Pearl</a>
        </Navbar.Brand>
        <Navbar.Toggle />
      </Navbar.Header>
      <Navbar.Collapse>
        <Nav>
          <NavItem href='/#/movies'>Movies</NavItem>
	  <NavItem href='/#'>TV</NavItem>
	</Nav>
      </Navbar.Collapse>
    </Navbar>
    )
  },
});

module.exports = NavigationContainer;
