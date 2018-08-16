const path = require('path')
module.exports = {
  siteMetadata: {
    title: 'Shorted',
  },
  plugins: [
    'gatsby-plugin-react-helmet',
    'gatsby-plugin-styled-components',
    `gatsby-transformer-yaml`,
    'gatsby-plugin-antd',
    {
      resolve: `gatsby-source-filesystem`,
      options: {
        path: `${__dirname}/data/stocks`,
        name: 'stocks',
      },
    }
  ],
  
}
