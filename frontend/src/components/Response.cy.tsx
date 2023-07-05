import React from 'react'
import Response from './Response'

describe('<Response />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Response />)
  })
})