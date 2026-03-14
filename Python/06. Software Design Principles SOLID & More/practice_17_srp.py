"""
Practice 17: Applying SRP (Single Responsibility Principle)
Task: Refactor OnDemandAgentService so each class has a single responsibility.

How to run:
  python practice_17_srp.py
"""

from dataclasses import dataclass
from datetime import datetime


class AssignmentRepository:
    """Persists assignment records."""

    def save(self, assignment: dict) -> None:
        # In a real system this would write to a database.
        print(f"[DB] Assignment saved: {assignment}")


class NotificationService:
    """Sends notifications to customers and agents."""

    def notify_customer(self, customer_name: str, agent_name: str) -> None:
        print(f"[Notify] Dear {customer_name}, {agent_name} is on the way.")

    def notify_agent(self, agent_name: str, customer_name: str) -> None:
        print(f"[Notify] Hi {agent_name}, you have a new trip for {customer_name}.")


class AuditLogger:
    """Writes domain audit logs."""

    def log_assignment(self, customer_name: str, agent_name: str) -> None:
        print(f"[Audit] {datetime.now().isoformat()} Assigned {agent_name} to {customer_name}")


@dataclass
class Customer:
    name: str


@dataclass
class Agent:
    name: str
    is_available: bool = True


class AgentMatcher:
    """Selects an available agent."""

    def find_available_agent(self, agents: list[Agent]) -> Agent:
        for agent in agents:
            if agent.is_available:
                return agent
        raise RuntimeError("No available agent found.")


class OnDemandAgentService:
    """Coordinates assignment flow only (single orchestration responsibility)."""

    def __init__(
        self,
        matcher: AgentMatcher,
        repository: AssignmentRepository,
        notifier: NotificationService,
        logger: AuditLogger,
    ):
        self.matcher = matcher
        self.repository = repository
        self.notifier = notifier
        self.logger = logger

    def assign_agent(self, customer: Customer, agents: list[Agent]) -> Agent:
        agent = self.matcher.find_available_agent(agents)
        agent.is_available = False

        assignment = {
            "customer": customer.name,
            "agent": agent.name,
            "assigned_at": datetime.now().isoformat(),
        }
        self.repository.save(assignment)
        self.notifier.notify_customer(customer.name, agent.name)
        self.notifier.notify_agent(agent.name, customer.name)
        self.logger.log_assignment(customer.name, agent.name)
        return agent


def main() -> None:
    customer = Customer("Afsana")
    agents = [Agent("Robin", is_available=False), Agent("Pulok", is_available=True)]

    service = OnDemandAgentService(
        matcher=AgentMatcher(),
        repository=AssignmentRepository(),
        notifier=NotificationService(),
        logger=AuditLogger(),
    )

    assigned = service.assign_agent(customer, agents)
    print(f"Result: {assigned.name} assigned successfully.")


if __name__ == "__main__":
    main()
