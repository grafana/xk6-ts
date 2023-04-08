import joe from "./joe.json"
import {NewUser} from "./user.ts"

export default function () {
  console.log(joe)
  console.log(NewUser("jim", 33))
}
