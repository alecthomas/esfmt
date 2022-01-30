<div>
  <h1>Hello</h1>
</div>;
<div>
  <h1>Hello</h1>
  World
</div>;
const CustomComp = (props) => <div>{props.children}</div>
<CustomComp>
  <div>Hello World</div>
  {"This is just a JS expression..." + 1000}
</CustomComp>
