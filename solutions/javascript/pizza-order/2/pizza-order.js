/// <reference path="./global.d.ts" />
//
// @ts-check

/**
 * Determine the price of the pizza given the pizza and optional extras
 *
 * @param {Pizza} pizza name of the pizza to be made
 * @param {Extra[]} extras list of extras
 *
 * @returns {number} the price of the pizza
 */
export function pizzaPrice(pizza, ...extras) {
  let price = 0;
  let stuffList = [pizza, ...extras];
  let getPrice = (stuff) => {
    switch(stuff) {
      case 'Margherita':
        return 7;
      case 'Caprese':
        return 9;
      case 'Formaggio':
        return 10;
      case 'ExtraSauce':
        return 1;
      case 'ExtraToppings':
        return 2;
      default:
        return 0;  
    }
  }
  for (let stuff of stuffList) {
    price += getPrice(stuff);
  }
  return price;
}

/**
 * Calculate the price of the total order, given individual orders
 *
 * (HINT: For this exercise, you can take a look at the supplied "global.d.ts" file
 * for a more info about the type definitions used)
 *
 * @param {PizzaOrder[]} pizzaOrders a list of pizza orders
 * @returns {number} the price of the total order
 */
export function orderPrice(pizzaOrders) {
  let total = 0;
  pizzaOrders.forEach(pizza => { 
    total += pizzaPrice(pizza.pizza, ...pizza.extras);
  });
  return total;
}
