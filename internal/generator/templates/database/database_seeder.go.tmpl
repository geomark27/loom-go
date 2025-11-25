package seeders

import (
	"log"

	"gorm.io/gorm"
)

// DatabaseSeeder orchestrates all seeders
type DatabaseSeeder struct{}

// Run executes all registered seeders
func (s *DatabaseSeeder) Run(db *gorm.DB) error {
	log.Println("🌱 Running seeders...")

	for _, seeder := range AllSeeders {
		if err := seeder.Run(db); err != nil {
			log.Printf("❌ Seeder failed: %v", err)
			return err
		}
	}

	log.Println("✅ All seeders executed successfully")
	return nil
}
