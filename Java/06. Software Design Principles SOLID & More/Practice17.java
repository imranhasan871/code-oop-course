/**
 * Practice 17: Applying SRP (Single Responsibility Principle)
 * Task: Refactor OnDemandAgentService so each class has a single responsibility.
 *
 * How to compile and run:
 *   javac Practice17.java
 *   java Practice17
 */

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

public class Practice17 {

    static class Customer {
        String name;

        Customer(String name) {
            this.name = name;
        }
    }

    static class Agent {
        String name;
        boolean isAvailable;

        Agent(String name, boolean isAvailable) {
            this.name = name;
            this.isAvailable = isAvailable;
        }
    }

    static class AssignmentRepository {
        void save(String customerName, String agentName) {
            System.out.println("[DB] Assignment saved: customer=" + customerName + ", agent=" + agentName);
        }
    }

    static class NotificationService {
        void notifyCustomer(String customerName, String agentName) {
            System.out.println("[Notify] Dear " + customerName + ", " + agentName + " is on the way.");
        }

        void notifyAgent(String agentName, String customerName) {
            System.out.println("[Notify] Hi " + agentName + ", new trip assigned for " + customerName + ".");
        }
    }

    static class AuditLogger {
        void logAssignment(String customerName, String agentName) {
            System.out.println("[Audit] " + LocalDateTime.now() + " Assigned " + agentName + " to " + customerName);
        }
    }

    static class AgentMatcher {
        Agent findAvailableAgent(List<Agent> agents) {
            for (Agent agent : agents) {
                if (agent.isAvailable) {
                    return agent;
                }
            }
            throw new IllegalStateException("No available agent found.");
        }
    }

    static class OnDemandAgentService {
        private final AgentMatcher matcher;
        private final AssignmentRepository repository;
        private final NotificationService notifier;
        private final AuditLogger logger;

        OnDemandAgentService(
            AgentMatcher matcher,
            AssignmentRepository repository,
            NotificationService notifier,
            AuditLogger logger
        ) {
            this.matcher = matcher;
            this.repository = repository;
            this.notifier = notifier;
            this.logger = logger;
        }

        Agent assignAgent(Customer customer, List<Agent> agents) {
            Agent agent = matcher.findAvailableAgent(agents);
            agent.isAvailable = false;

            repository.save(customer.name, agent.name);
            notifier.notifyCustomer(customer.name, agent.name);
            notifier.notifyAgent(agent.name, customer.name);
            logger.logAssignment(customer.name, agent.name);

            return agent;
        }
    }

    public static void main(String[] args) {
        Customer customer = new Customer("Afsana");
        List<Agent> agents = new ArrayList<>();
        agents.add(new Agent("Robin", false));
        agents.add(new Agent("Pulok", true));

        OnDemandAgentService service = new OnDemandAgentService(
            new AgentMatcher(),
            new AssignmentRepository(),
            new NotificationService(),
            new AuditLogger()
        );

        Agent assigned = service.assignAgent(customer, agents);
        System.out.println("Result: " + assigned.name + " assigned successfully.");
    }
}
