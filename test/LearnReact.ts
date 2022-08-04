import { Given, When, Then, DataTable } from '@badeball/cypress-cucumber-preprocessor'

Given('the user has browsed to the login page', () => {
  // TODO: implement step
  console.log('the user has browsed to the login page')
});

Given('a user {string} has been created with the following details:', (a: string, dt: DataTable) => {
  // TODO: implement step
  console.log('a user "asdfasd" has been created with the following details:', a, dt)
});

When('the user enters the following details in the login form:', (dt: DataTable) => {
  // TODO: implement step
  console.log('the user enters the following details in the login form:', dt)
});

When('the user logs in', () => {
  // TODO: implement step
  console.log('the user logs in')
});

Then('the user {int} be redirected to the homepage', (a: number, ds: string) => {
  // TODO: implement step
  console.log('the user 99 be redirected to the homepage', a, ds)
});
