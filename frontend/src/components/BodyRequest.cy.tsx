import React from 'react'
import BodyRequest from './BodyRequest'

describe('<BodyRequest />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<BodyRequest />)
  })
})