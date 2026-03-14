/**
 * Practice 17: Applying SRP (Single Responsibility Principle)
 * Task: Refactor OnDemandAgentService so each class has a single responsibility.
 *
 * How to run (using dotnet-script):
 *   dotnet script Practice17.cs
 */

using System;
using System.Collections.Generic;

class Practice17
{
    class Customer
    {
        public string Name;

        public Customer(string name)
        {
            Name = name;
        }
    }

    class Agent
    {
        public string Name;
        public bool IsAvailable;

        public Agent(string name, bool isAvailable)
        {
            Name = name;
            IsAvailable = isAvailable;
        }
    }

    class AssignmentRepository
    {
        public void Save(string customerName, string agentName)
        {
            Console.WriteLine($"[DB] Assignment saved: customer={customerName}, agent={agentName}");
        }
    }

    class NotificationService
    {
        public void NotifyCustomer(string customerName, string agentName)
        {
            Console.WriteLine($"[Notify] Dear {customerName}, {agentName} is on the way.");
        }

        public void NotifyAgent(string agentName, string customerName)
        {
            Console.WriteLine($"[Notify] Hi {agentName}, new trip assigned for {customerName}.");
        }
    }

    class AuditLogger
    {
        public void LogAssignment(string customerName, string agentName)
        {
            Console.WriteLine($"[Audit] {DateTime.Now:O} Assigned {agentName} to {customerName}");
        }
    }

    class AgentMatcher
    {
        public Agent FindAvailableAgent(List<Agent> agents)
        {
            foreach (var agent in agents)
            {
                if (agent.IsAvailable)
                    return agent;
            }
            throw new InvalidOperationException("No available agent found.");
        }
    }

    class OnDemandAgentService
    {
        private readonly AgentMatcher matcher;
        private readonly AssignmentRepository repository;
        private readonly NotificationService notifier;
        private readonly AuditLogger logger;

        public OnDemandAgentService(
            AgentMatcher matcher,
            AssignmentRepository repository,
            NotificationService notifier,
            AuditLogger logger)
        {
            this.matcher = matcher;
            this.repository = repository;
            this.notifier = notifier;
            this.logger = logger;
        }

        public Agent AssignAgent(Customer customer, List<Agent> agents)
        {
            var agent = matcher.FindAvailableAgent(agents);
            agent.IsAvailable = false;

            repository.Save(customer.Name, agent.Name);
            notifier.NotifyCustomer(customer.Name, agent.Name);
            notifier.NotifyAgent(agent.Name, customer.Name);
            logger.LogAssignment(customer.Name, agent.Name);

            return agent;
        }
    }

    static void Main(string[] args)
    {
        var customer = new Customer("Afsana");
        var agents = new List<Agent>
        {
            new Agent("Robin", false),
            new Agent("Pulok", true)
        };

        var service = new OnDemandAgentService(
            new AgentMatcher(),
            new AssignmentRepository(),
            new NotificationService(),
            new AuditLogger());

        var assigned = service.AssignAgent(customer, agents);
        Console.WriteLine($"Result: {assigned.Name} assigned successfully.");
    }
}
