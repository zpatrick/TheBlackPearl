var axios = require('axios');

axios.defaults.baseURL = 'http://zp-dev-1-1236183437.us-west-1.elb.amazonaws.com:8000'

api = {
  listMovies: function() {
    return axios.get('/movies').
    then(function(response) {
      return response.data;
    });
  },
}

module.exports = api
