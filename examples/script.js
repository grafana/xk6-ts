import { newUser } from "./user";

export default () => {
  const user = newUser("John");
  console.log(user);
};
