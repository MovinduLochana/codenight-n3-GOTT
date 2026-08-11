import { integer, text, boolean, pgTable } from "drizzle-orm/pg-core";

// just copied fomnr dcos
export const todo = pgTable("todo", {
  id: integer("id").primaryKey(),
  text: text("text").notNull(),
  done: boolean("done").default(false).notNull(),
});
