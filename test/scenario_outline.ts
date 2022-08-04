import { Given, When, Then, DataTable } from '@badeball/cypress-cucumber-preprocessor'

Given('there are {int} cucumbers', (a: number) => {
  // TODO: implement step
  console.log('there are 27 cucumbers', a)
});

Given('there word hello', () => {
  // TODO: implement step
  console.log('there word hello')
});

Given('there word cruel', () => {
  // TODO: implement step
  console.log('there word cruel')
});

Given('there word world', () => {
  // TODO: implement step
  console.log('there word world')
});

When('I remove letter l', () => {
  // TODO: implement step
  console.log('I remove letter l')
});

When('I remove letter w', () => {
  // TODO: implement step
  console.log('I remove letter w')
});

When('I eat {int} cucumbers', (a: number) => {
  // TODO: implement step
  console.log('I eat 7 cucumbers', a)
});

When('I remove letter e', () => {
  // TODO: implement step
  console.log('I remove letter e')
});

Then('I should have {int} cucumbers', (a: number) => {
  // TODO: implement step
  console.log('I should have 20 cucumbers', a)
});

Then('I should have word hllo', () => {
  // TODO: implement step
  console.log('I should have word hllo')
});

Then('I should have word crue', () => {
  // TODO: implement step
  console.log('I should have word crue')
});

Then('I should have word orld', () => {
  // TODO: implement step
  console.log('I should have word orld')
});
