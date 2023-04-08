interface User {
  name: string
  age: number
}

function NewUser(name: string, age: number) : User {
  return {
    name: name,
    age: age
  }
}

export {NewUser}
