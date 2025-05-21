"use client";

import { graphqlClient } from "@/lib/graphql";
import { TodoResponse } from "@/lib/types";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { Todo } from "@/lib/types";

const GET_TODO = `
  query GetTodo($id: ID!) {
    todo(id: $id) {
      id
      text
      done
      user {
        name
      }
    }
  }
`;

export default function TodoPage() {
  const { id } = useParams();

  const [todo, setTodo] = useState<Todo | null>(null);

  useEffect(() => {
    (async () => {
      const { todo } = await graphqlClient.request<TodoResponse>(GET_TODO, { id });
      setTodo(todo);
    })();
  }, [id]);

  if (!todo) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>{todo.text}</h1>
      <p>{todo.done ? "Done" : "Not done"}</p>
      <p>by {todo.user.name}</p>
    </div>
  );
}