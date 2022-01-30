interface ClickableProps {
  children: JSX.Element[] | JSX.Element;
}
 
interface HomeProps extends ClickableProps {
  home: JSX.Element;
}
 
interface SideProps extends ClickableProps {
  side: JSX.Element | string;
}
 
function MainButton(prop: HomeProps): JSX.Element;
function MainButton(prop: SideProps): JSX.Element;
function MainButton(prop: ClickableProps): JSX.Element {
  // ...
}
