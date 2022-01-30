class MyComponent {
  render() {}
}
// use a construct signature
const myComponent = new MyComponent();
// element class type => MyComponent
// element instance type => { render: () => void }
function MyFactoryFunction() {
  return {
    render: () => {},
  };
}
// use a call signature
const myComponent = MyFactoryFunction();
// element class type => MyFactoryFunction
// element instance type => { render: () => void }
