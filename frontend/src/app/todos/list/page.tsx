import { graphqlClient } from "@/lib/graphql";
import { TodosResponse } from "@/lib/types";
import Link from "next/link";
const GET_TODOS = `
  query {
    todos {
      id
      text
      done
      user {
        id
        name
      }
    }
  }
`;

export default async function ListPage() {
  const { todos } = await graphqlClient.request<TodosResponse>(GET_TODOS);
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <h1>Todos</h1>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            <Link href={`/todos/${todo.id}`}>{todo.text}</Link> - by{" "}
            {todo.user.name}
          </li>
        ))}
      </ul>
    </div>
  );
}
