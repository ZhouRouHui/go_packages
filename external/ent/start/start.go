package start

import (
	"context"
	"fmt"
	"go_packages/external/ent/ent"
	"go_packages/external/ent/ent/car"
	"go_packages/external/ent/ent/group"
	"go_packages/external/ent/ent/user"
	"log"
	"time"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// only 在找不到用户或找到多于一个用户时报错
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// create a new car with model "Tesla"
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// create a new car with model "Ford"
	ford, err := client.Car.
		Create().
		SetModel("ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddCars(tesla, ford).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// what about filtering specific cars.
	ford, err := a8m.QueryCars().Where(car.Model("ford")).Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// first, create the users.
	a8m, err := client.User.Create().SetAge(30).SetName("Ariel").Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
	if err != nil {
		return err
	}
	// then, create the cars, and attach them to the users in the creation.
	err = client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).SetOwner(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.Create().SetModel("Mazda").SetRegisteredAt(time.Now()).SetOwner(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).SetOwner(neta).Exec(ctx)
	if err != nil {
		return err
	}

	// create the groups, and add their users in the creation
	err = client.Group.Create().SetName("gitlab").AddUsers(a8m, neta).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.Create().SetName("github").AddUsers(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("the graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.Query().Where(group.Name("github")).QueryUsers().QueryCars().All(ctx)
	if err != nil {
		return err
	}
	log.Println("cars returned: ", cars)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	a8m := client.User.Query().Where(
		user.HasCars(),
		user.Name("Ariel"),
	).
		OnlyX(ctx)
	cars, err := a8m.QueryGroups().QueryUsers().QueryCars().Where(car.Not(car.Model("Mazda"))).All(ctx)
	if err != nil {
		return err
	}
	log.Println("cars returned: ", cars)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		return err
	}
	log.Println("groups returned: ", groups)
	return nil
}
