package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type ContactManager struct {
	contacts []Contact
	nextID   int
}

func (cm *ContactManager) AddContact(name, email, phone string) Contact {
	contact := Contact{
		ID:    cm.nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	cm.contacts = append(cm.contacts, contact)
	cm.nextID++
	return contact
}

func (cm *ContactManager) GetContact(id int) (Contact, error) {
	for _, c := range cm.contacts {
		if c.ID == id {
			return c, nil
		}
	}
	return Contact{}, fmt.Errorf("contact avec ID %d introuvable", id)
}

func (cm *ContactManager) ListContacts() []Contact {
	return cm.contacts
}

func (cm *ContactManager) DeleteContact(id int) error {
	for i, c := range cm.contacts {
		if c.ID == id {
			cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact avec ID %d introuvable", id)
}

func (cm *ContactManager) SearchByName(query string) []Contact {
	var results []Contact
	query = strings.ToLower(query)
	for _, c := range cm.contacts {
		if strings.Contains(strings.ToLower(c.Name), query) {
			results = append(results, c)
		}
	}
	return results
}

func (cm *ContactManager) SearchByEmail(email string) (Contact, error) {
	for _, c := range cm.contacts {
		if strings.EqualFold(c.Email, email) {
			return c, nil
		}
	}
	return Contact{}, fmt.Errorf("aucun contact avec l'email %s", email)
}

func main() {
	cm := &ContactManager{}

	cm.AddContact("Alice Dupont", "alice@email.com", "0601020304")
	cm.AddContact("Bob Martin", "bob@email.com", "0605060708")

	for {
		fmt.Println("\n=== Gestionnaire de Contacts ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Rechercher un contact par ID")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Rechercher par nom")
		fmt.Println("6. Rechercher par email")
		fmt.Println("7. Quitter")
		fmt.Print("\nChoix: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email, phone string
			fmt.Print("Nom: ")
			fmt.Scan(&name)
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Téléphone: ")
			fmt.Scan(&phone)

			contact := cm.AddContact(name, email, phone)
			fmt.Printf(":) Contact ajouté (ID: %d)\n", contact.ID)

		case 2:
			contacts := cm.ListContacts()
			if len(contacts) == 0 {
				fmt.Println("Aucun contact.")
			} else {
				fmt.Println("\nListe des contacts:")
				for _, c := range contacts {
					fmt.Printf("  [%d] %s - %s - %s\n", c.ID, c.Name, c.Email, c.Phone)
				}
			}

		case 3:
			var id int
			fmt.Print("ID du contact: ")
			fmt.Scan(&id)

			contact, err := cm.GetContact(id)
			if err != nil {
				fmt.Printf("X %s\n", err)
			} else {
				fmt.Printf("\n[%d] %s\nEmail: %s\nTél: %s\n",
					contact.ID, contact.Name, contact.Email, contact.Phone)
			}

		case 4:
			var id int
			fmt.Print("ID du contact à supprimer: ")
			fmt.Scan(&id)

			err := cm.DeleteContact(id)
			if err != nil {
				fmt.Printf("X %s\n", err)
			} else {
				fmt.Println(":) Contact supprimé")
			}

		case 5:
			var query string
			fmt.Print("Nom à rechercher: ")
			fmt.Scan(&query)

			results := cm.SearchByName(query)
			if len(results) == 0 {
				fmt.Println("Aucun résultat.")
			} else {
				fmt.Printf("\n%d résultat(s):\n", len(results))
				for _, c := range results {
					fmt.Printf("  [%d] %s - %s - %s\n", c.ID, c.Name, c.Email, c.Phone)
				}
			}

		case 6:
			var email string
			fmt.Print("Email à rechercher: ")
			fmt.Scan(&email)

			contact, err := cm.SearchByEmail(email)
			if err != nil {
				fmt.Printf("X %s\n", err)
			} else {
				fmt.Printf("\n[%d] %s\nEmail: %s\nTél: %s\n",
					contact.ID, contact.Name, contact.Email, contact.Phone)
			}

		case 7:
			fmt.Println("Au revoir!")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
