const hostname = window && window.location && window.location.hostname;

const devEndpoints = {
  api: 'http://localhost:9999/graphql',
};

const prodEndpoint = {
  api: '/graphql',
};

const Endpoints = () => {
  switch (hostname) {
    case 'arijeetbaruah.herokuapp.com': {
      return prodEndpoint;
    }
    default: {
      return devEndpoints;
    }
  }
};

export default Endpoints();
