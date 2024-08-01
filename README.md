type User struct{
int $age
varchar(10) $name
}

type Employee struct{
User $user
int $salary
}

Employee $theBest

$theBest.user.age=27
$theBest.user.name='yigit'
