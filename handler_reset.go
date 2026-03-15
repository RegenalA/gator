package main

import (
        "context"
        "fmt"
)

func handlerReset(s *state, cmd command) error {
        err := s.db.DeleteUser(context.Background())
        if err != nil {
                return fmt.Errorf("couldn't delete user: %w", err)
        }
        fmt.Println("User deleted successfully:")
        return nil
}
