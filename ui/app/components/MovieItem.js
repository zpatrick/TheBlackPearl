var React = require('react');
var PropTypes = React.PropTypes;
var Bootstrap = require('react-bootstrap');
var Col = Bootstrap.Col;
var Image = Bootstrap.Image;

var MovieItem = React.createClass({
  PropTypes: {
    title: PropTypes.string.isRequired,
    description: PropTypes.string,
    posterURL: PropTypes.string,
  },
  render: function() {
    return (
      <Col md={4}>
        <a href='/#'>
	  <Image src={this.props.posterURL} responsive />
	</a>
	<h3>
	  <a href='/#'>{this.props.title}</a>
	</h3>
	<p>{this.props.description}</p>
      </Col>
    )
  }
});

module.exports = MovieItem;
