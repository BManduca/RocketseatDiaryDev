interface User {
    name: string, 
    bio: string,
    age: number
}

function sumAge(users: User[]) {
    let sum = 0;

    for(const user of users) {
        sum += user.age;
    }

    return sum;
}

const sumAllUsersAges = sumAge([
    {
        name: 'Brunno',
        bio: 'Desenvolvedor',
        age: 31
    }
])

console.log('O Total das idades: ' + sumAllUsersAges)