package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// -------------------- Structures --------------------

type Bank struct {
	mu        sync.RWMutex
	Name      string
	bankMoney float64
	Deposit   float64
	Credit    float64
	Clients   map[string]*Client
}

type ClientType int

const (
	DepositClient ClientType = iota
	CreditClient
)

// Client представляє клієнта банку.
type Client struct {
	mu             sync.RWMutex
	Name           string
	Surname        string
	AccountNumber  string
	cDeposit       float64
	cCredit        float64
	cash           float64
	Type           ClientType
	creditLimit    float64
	depositLimit   float64
	everTookCredit bool
}

// -------------------- Constructors --------------------

func NewBank(name string, initialMoney float64) *Bank {
	return &Bank{
		Name:      name,
		bankMoney: initialMoney,
		Deposit:   0,
		Credit:    0,
		Clients:   make(map[string]*Client),
	}
}

func NewDepositClient(name, surname, account string, initialCash, depositLimit float64) *Client {
	return &Client{
		Name:          name,
		Surname:       surname,
		AccountNumber: account,
		cDeposit:      0,
		cCredit:       0,
		cash:          initialCash,
		Type:          DepositClient,
		depositLimit:  depositLimit,
	}
}

func NewCreditClient(name, surname, account string, initialCash, creditLimit float64) *Client {
	return &Client{
		Name:          name,
		Surname:       surname,
		AccountNumber: account,
		cDeposit:      0,
		cCredit:       0,
		cash:          initialCash,
		Type:          CreditClient,
		creditLimit:   creditLimit,
	}
}

// -------------------- Bank methods (set/get & actions) --------------------

func (b *Bank) GetName() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Name
}

func (b *Bank) GetBankMoney() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.bankMoney
}

func (b *Bank) GetDepositSum() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Deposit
}

func (b *Bank) GetCreditSum() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Credit
}

// AddClient додає клієнта до банку
func (b *Bank) AddClient(c *Client) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.Clients[c.AccountNumber]; ok {
		return fmt.Errorf("account %s already exists", c.AccountNumber)
	}
	b.Clients[c.AccountNumber] = c
	return nil
}

// GetClientBySurname повертає перший клієнт з таким прізвищем
func (b *Bank) GetClientBySurname(surname string) (*Client, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, c := range b.Clients {
		c.mu.RLock()
		if strings.EqualFold(c.Surname, surname) {
			c.mu.RUnlock()
			return c, true
		}
		c.mu.RUnlock()
	}
	return nil, false
}

func (b *Bank) GetClientByAccount(acc string) (*Client, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	c, ok := b.Clients[acc]
	return c, ok
}

// Bank отримує депозит від клієнта (client -> bank)
func (b *Bank) ReceiveDeposit(client *Client, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be > 0")
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	if client.cash < amount {
		return fmt.Errorf("client %s %s doesn't have enough cash to deposit", client.Name, client.Surname)
	}
	if client.cDeposit+amount > client.depositLimit {
		return fmt.Errorf("deposit would exceed client's deposit limit")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client.cash -= amount
	client.cDeposit += amount

	b.Deposit += amount
	b.bankMoney += amount
	return nil
}

// Bank повертає депозит клієнту (bank -> client), тобто знімає депозит
func (b *Bank) WithdrawDeposit(client *Client, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be > 0")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	client.mu.Lock()
	defer client.mu.Unlock()

	if client.cDeposit < amount {
		return fmt.Errorf("client deposit balance too small")
	}
	if b.bankMoney < amount {
		return fmt.Errorf("bank doesn't have enough money to return deposit right now")
	}

	client.cDeposit -= amount
	client.cash += amount

	b.Deposit -= amount
	b.bankMoney -= amount
	return nil
}

// IssueCredit — банк видає кредит клієнту (bank -> client), client.cCredit збільшується
func (b *Bank) IssueCredit(client *Client, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be > 0")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client.mu.Lock()
	defer client.mu.Unlock()

	if client.cCredit+amount > client.creditLimit {
		return fmt.Errorf("would exceed client's credit limit")
	}
	if b.bankMoney < amount {
		return fmt.Errorf("bank doesn't have enough money to issue credit")
	}

	client.cCredit += amount
	client.cash += amount
	client.everTookCredit = true

	b.Credit += amount
	b.bankMoney -= amount
	return nil
}

// RepayCredit — клієнт повертає кредит банку (client -> bank)
func (b *Bank) RepayCredit(client *Client, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be > 0")
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	if client.cash < amount {
		return fmt.Errorf("client doesn't have enough cash to repay")
	}
	if client.cCredit <= 0 {
		return fmt.Errorf("client has no credit to repay")
	}
	if amount > client.cCredit {
		amount = client.cCredit
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	client.cash -= amount
	client.cCredit -= amount

	b.Credit -= amount
	b.bankMoney += amount
	return nil
}

// -------------------- Client getters/setters --------------------

func (c *Client) GetFullName() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("%s %s", c.Name, c.Surname)
}

func (c *Client) GetAccountNumber() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.AccountNumber
}

func (c *Client) GetDepositBalance() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cDeposit
}

func (c *Client) GetCreditBalance() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cCredit
}

func (c *Client) GetCash() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cash
}

// -------------------- Bot behavior --------------------

// StartClientBot запускає горутину, яка кожну секунду робить випадкову операцію.
// Контекст використовується для відміни; wg для очікування завершення всіх ботів.
func StartClientBot(ctx context.Context, wg *sync.WaitGroup, bank *Bank, client *Client) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(len(client.AccountNumber))))
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("[Bot %s] stopped by context\n", client.GetAccountNumber())
				return
			case <-t.C:
				switch client.Type {
				case DepositClient:
					action := r.Intn(2)
					if action == 0 {
						client.mu.RLock()
						max := client.cash
						limitLeft := client.depositLimit - client.cDeposit
						client.mu.RUnlock()
						if max <= 0 || limitLeft <= 0 {
							continue
						}
						amt := 1 + r.Float64()*(minFloat64(max, limitLeft)-1)
						if amt < 0.01 {
							amt = minFloat64(max, limitLeft)
						}
						err := bank.ReceiveDeposit(client, round2(amt))
						if err != nil {
						} else {
							fmt.Printf("[DepositBot %s] deposited %.2f -> deposit balance: %.2f, cash: %.2f\n",
								client.AccountNumber, round2(amt), client.GetDepositBalance(), client.GetCash())
						}
					} else {
						client.mu.RLock()
						maxWithdraw := client.cDeposit
						client.mu.RUnlock()
						if maxWithdraw <= 0 {
							continue
						}
						amt := 1 + r.Float64()*(maxWithdraw-1)
						if amt < 0.01 {
							amt = maxWithdraw
						}
						err := bank.WithdrawDeposit(client, round2(amt))
						if err != nil {
						} else {
							fmt.Printf("[DepositBot %s] withdrew %.2f <- deposit left: %.2f, cash: %.2f\n",
								client.AccountNumber, round2(amt), client.GetDepositBalance(), client.GetCash())
						}
						if client.GetDepositBalance() == 0 {
							fmt.Printf("[DepositBot %s] вибрав увесь ліміт депозиту (депозит = 0). Бот завершує роботу.\n", client.AccountNumber)
							return
						}
					}
				case CreditClient:
					action := r.Intn(2)
					if action == 0 {
						client.mu.RLock()
						limitLeft := client.creditLimit - client.cCredit
						client.mu.RUnlock()
						if limitLeft <= 0 {
							continue
						}
						bank.mu.RLock()
						bankHas := bank.bankMoney
						bank.mu.RUnlock()
						max := minFloat64(limitLeft, bankHas)
						if max <= 0 {
							continue
						}
						amt := 1 + r.Float64()*(max-1)
						if amt < 0.01 {
							amt = max
						}
						err := bank.IssueCredit(client, round2(amt))
						if err != nil {

						} else {
							fmt.Printf("[CreditBot %s] took credit %.2f -> credit balance: %.2f, cash: %.2f\n",
								client.AccountNumber, round2(amt), client.GetCreditBalance(), client.GetCash())
						}
					} else {
						client.mu.RLock()
						maxRepay := minFloat64(client.cash, client.cCredit)
						client.mu.RUnlock()
						if maxRepay <= 0 {
							continue
						}
						amt := 1 + r.Float64()*(maxRepay-1)
						if amt < 0.01 {
							amt = maxRepay
						}
						err := bank.RepayCredit(client, round2(amt))
						if err != nil {

						} else {
							fmt.Printf("[CreditBot %s] repaid %.2f -> credit left: %.2f, cash: %.2f\n",
								client.AccountNumber, round2(amt), client.GetCreditBalance(), client.GetCash())
						}
						if client.GetCreditBalance() == 0 && client.everTookCredit {
							fmt.Printf("[CreditBot %s] повернув кредит повністю. Бот завершує роботу.\n", client.AccountNumber)
							return
						}
					}
				}
			}
		}
	}()
}

// -------------------- Helpers --------------------

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func round2(v float64) float64 {
	return float64(int(v*100)) / 100.0
}

// -------------------- Console UI --------------------

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	var bank *Bank
	var wg sync.WaitGroup
	var botsCtx context.Context
	var botsCancel context.CancelFunc
	botsCtx, botsCancel = context.WithCancel(context.Background())

mainloop:
	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1) Створення банку")
		fmt.Println("2) Створення клієнта для кредитів")
		fmt.Println("3) Створення клієнта для депозитів")
		fmt.Println("4) Виведення інформації про клієнта за прізвищем")
		fmt.Println("5) Виведення інформації про клієнта за номером рахунку")
		fmt.Println("6) Завершення")
		fmt.Print("Виберіть пункт: ")
		choiceRaw, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceRaw)

		switch choice {
		case "1":
			// Create bank
			fmt.Print("Введіть назву банку: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if name == "" {
				fmt.Println("Назва не може бути порожньою")
				continue
			}
			fmt.Print("Початкові власні кошти банку (число): ")
			var initial float64
			_, err := fmt.Fscanln(reader, &initial)
			if err != nil {
				line, _ := reader.ReadString('\n')
				fmt.Sscan(strings.TrimSpace(line), &initial)
			}
			if initial < 0 {
				initial = 0
			}
			botsCancel()
			wg.Wait()
			botsCtx, botsCancel = context.WithCancel(context.Background())

			bank = NewBank(name, initial)
			fmt.Printf("Банк %s створено з власними коштами: %.2f\n", bank.Name, bank.GetBankMoney())

		case "2":
			if bank == nil {
				fmt.Println("Створіть спочатку банк (пункт 1).")
				continue
			}
			// Create credit client
			fmt.Print("Ім'я: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Прізвище: ")
			surname, _ := reader.ReadString('\n')
			surname = strings.TrimSpace(surname)
			fmt.Print("Номер рахунку (унікальний): ")
			account, _ := reader.ReadString('\n')
			account = strings.TrimSpace(account)
			fmt.Print("Початкова готівка клієнта (cash): ")
			var cash float64
			_, err := fmt.Fscanln(reader, &cash)
			if err != nil {
				line, _ := reader.ReadString('\n')
				fmt.Sscan(strings.TrimSpace(line), &cash)
			}
			fmt.Print("Ліміт кредиту: ")
			var limit float64
			_, err2 := fmt.Fscanln(reader, &limit)
			if err2 != nil {
				line, _ := reader.ReadString('\n')
				fmt.Sscan(strings.TrimSpace(line), &limit)
			}
			if cash < 0 {
				cash = 0
			}
			if limit <= 0 {
				limit = 1000
			}
			client := NewCreditClient(name, surname, account, cash, limit)
			if err := bank.AddClient(client); err != nil {
				fmt.Println("Помилка додавання клієнта:", err)
			} else {
				fmt.Printf("Клієнт %s створений (credit client). Запуск бота...\n", client.GetAccountNumber())
				StartClientBot(botsCtx, &wg, bank, client)
			}

		case "3":
			if bank == nil {
				fmt.Println("Створіть спочатку банк (пункт 1).")
				continue
			}
			// Create deposit client
			fmt.Print("Ім'я: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Прізвище: ")
			surname, _ := reader.ReadString('\n')
			surname = strings.TrimSpace(surname)
			fmt.Print("Номер рахунку (унікальний): ")
			account, _ := reader.ReadString('\n')
			account = strings.TrimSpace(account)
			fmt.Print("Початкова готівка клієнта (cash): ")
			var cash float64
			_, err := fmt.Fscanln(reader, &cash)
			if err != nil {
				line, _ := reader.ReadString('\n')
				fmt.Sscan(strings.TrimSpace(line), &cash)
			}
			fmt.Print("Ліміт депозиту: ")
			var limit float64
			_, err2 := fmt.Fscanln(reader, &limit)
			if err2 != nil {
				line, _ := reader.ReadString('\n')
				fmt.Sscan(strings.TrimSpace(line), &limit)
			}
			if cash < 0 {
				cash = 0
			}
			if limit <= 0 {
				limit = 1000
			}
			client := NewDepositClient(name, surname, account, cash, limit)
			if err := bank.AddClient(client); err != nil {
				fmt.Println("Помилка додавання клієнта:", err)
			} else {
				fmt.Printf("Клієнт %s створений (deposit client). Запуск бота...\n", client.GetAccountNumber())
				StartClientBot(botsCtx, &wg, bank, client)
			}

		case "4":
			if bank == nil {
				fmt.Println("Створіть спочатку банк.")
				continue
			}
			fmt.Print("Введіть прізвище клієнта: ")
			surname, _ := reader.ReadString('\n')
			surname = strings.TrimSpace(surname)
			c, ok := bank.GetClientBySurname(surname)
			if !ok {
				fmt.Println("Клієнта не знайдено.")
				continue
			}
			printClientInfo(c)

		case "5":
			if bank == nil {
				fmt.Println("Створіть спочатку банк.")
				continue
			}
			fmt.Print("Введіть номер рахунку: ")
			acc, _ := reader.ReadString('\n')
			acc = strings.TrimSpace(acc)
			c, ok := bank.GetClientByAccount(acc)
			if !ok {
				fmt.Println("Клієнта не знайдено.")
				continue
			}
			printClientInfo(c)

		case "6":
			fmt.Println("Завершення програми...")
			botsCancel()
			wg.Wait()
			break mainloop

		default:
			fmt.Println("Невірний вибір")
		}
	}
}

// printClientInfo виводить актуальний стан клієнта
func printClientInfo(c *Client) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	typ := "DepositClient"
	if c.Type == CreditClient {
		typ = "CreditClient"
	}
	fmt.Printf("Клієнт: %s %s | Account: %s | Type: %s\n", c.Name, c.Surname, c.AccountNumber, typ)
	fmt.Printf("  Deposit balance: %.2f\n", c.cDeposit)
	fmt.Printf("  Credit balance:  %.2f\n", c.cCredit)
	fmt.Printf("  Cash:            %.2f\n", c.cash)
	fmt.Printf("  Deposit limit:   %.2f\n", c.depositLimit)
	fmt.Printf("  Credit limit:    %.2f\n", c.creditLimit)
}
