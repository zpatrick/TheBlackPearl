var React = require('react');
var Page = require('../components/Page');
var PageHeader = require('../components/PageHeader');
var MovieItem = require('../components/MovieItem');
var Bootstrap = require('react-bootstrap');
var Row = Bootstrap.Row;
var Col = Bootstrap.Col;
var Pagination = Bootstrap.Pagination;

var api = require('../utils/api');

const ROWS_PER_PAGE = 3;

var MovieContainer = React.createClass({
  getInitialState() {
    return {
      isLoading: true,
      movies: [],
      activePage: 1
    };
  },
  componentDidMount: function() {
    api.listMovies().
    then(function(movies) {
      this.setState({
        isLoading: false,
        movies: movies,
      });
    }.bind(this));
  },
  handlePageSelect(eventKey) {
    this.setState({
      activePage: eventKey
    });
  },
  getNumPages: function() {
    moviesPerPage = ROWS_PER_PAGE * 3
    return Math.ceil(this.state.movies.length / moviesPerPage );
  },
  getMovies: function() {
    numMovies = ROWS_PER_PAGE * 3
    start = (this.state.activePage - 1) * numMovies
    movies = this.state.movies.slice(start, start + numMovies);

    rows = [];
    for (row = 0, numRows = Math.ceil(movies.length / 3); row < numRows; row++) {
      i = row * 3;

      rows.push(
        <Row key={i}>
	  <MovieItem 
	    key={i}
	    title={movies[i].Title}
	    description={movies[i].Description}
	    posterURL={movies[i].PosterURL} />

	    { i+1 < movies.length ? 
	      <MovieItem
                key={i+1}
	        title={movies[i+1].Title}
	        description={movies[i+1].Description}
	        posterURL={movies[i+1].PosterURL} />
              :
	        null
	    }

            { i+2 < movies.length ?
	      <MovieItem
	        key={i+2}
	        title={movies[i+2].Title}
	        description={movies[i+2].Description}
	        posterURL={movies[i+2].PosterURL} />
	      :
	        null
              }
	    </Row>
      )
    }

    return rows;
  },
  render: function() {
    return (
      <Page>
        <Row>
          <Col lg={12}>
            <PageHeader text='Movies' subtext='' />
          </Col>
        </Row>
        <Pagination
	  prev
          next
	  ellipsis
	  boundaryLinks
	  items={this.getNumPages()}
          maxButtons={5}
          activePage={this.state.activePage}
          onSelect={this.handlePageSelect} />

        { this.state.isLoading ? <h1>Loading</h1> : this.getMovies()}
      </Page>
    )
  },
});

module.exports = MovieContainer;
