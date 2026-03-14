/**
 * Practice 17: Applying SRP (Single Responsibility Principle)
 * Task: Refactor OnDemandAgentService so each class has a single responsibility.
 *
 * How to run:
 *   go run practice-17.go
 */

package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Name string
}

type Agent struct {
	Name        string
	IsAvailable bool
}

type AssignmentRepository struct{}

func (r *AssignmentRepository) Save(customerName string, agentName string) {
	fmt.Printf("[DB] Assignment saved: customer=%s, agent=%s\n", customerName, agentName)
}

type NotificationService struct{}

func (n *NotificationService) NotifyCustomer(customerName string, agentName string) {
	fmt.Printf("[Notify] Dear %s, %s is on the way.\n", customerName, agentName)
}

func (n *NotificationService) NotifyAgent(agentName string, customerName string) {
	fmt.Printf("[Notify] Hi %s, new trip assigned for %s.\n", agentName, customerName)
}

type AuditLogger struct{}

func (l *AuditLogger) LogAssignment(customerName string, agentName string) {
	fmt.Printf("[Audit] %s Assigned %s to %s\n", time.Now().Format(time.RFC3339), agentName, customerName)
}

type AgentMatcher struct{}

func (m *AgentMatcher) FindAvailableAgent(agents []*Agent) (*Agent, error) {
	for _, agent := range agents {
		if agent.IsAvailable {
			return agent, nil
		}
	}
	return nil, fmt.Errorf("no available agent found")
}

type OnDemandAgentService struct {
	Matcher    *AgentMatcher
	Repository *AssignmentRepository
	Notifier   *NotificationService
	Logger     *AuditLogger
}

func (s *OnDemandAgentService) AssignAgent(customer *Customer, agents []*Agent) (*Agent, error) {
	agent, err := s.Matcher.FindAvailableAgent(agents)
	if err != nil {
		return nil, err
	}
	agent.IsAvailable = false

	s.Repository.Save(customer.Name, agent.Name)
	s.Notifier.NotifyCustomer(customer.Name, agent.Name)
	s.Notifier.NotifyAgent(agent.Name, customer.Name)
	s.Logger.LogAssignment(customer.Name, agent.Name)
	return agent, nil
}

func main() {
	customer := &Customer{Name: "Afsana"}
	agents := []*Agent{
		{Name: "Robin", IsAvailable: false},
		{Name: "Pulok", IsAvailable: true},
	}

	service := &OnDemandAgentService{
		Matcher:    &AgentMatcher{},
		Repository: &AssignmentRepository{},
		Notifier:   &NotificationService{},
		Logger:     &AuditLogger{},
	}

	assigned, err := service.AssignAgent(customer, agents)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %s assigned successfully.\n", assigned.Name)
}
