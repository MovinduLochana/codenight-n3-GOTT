import {
  boolean,
  integer,
  pgTable,
  text,
  timestamp,
} from "drizzle-orm/pg-core";

export const todo = pgTable("todo", {
  id: integer("id").primaryKey(),
  text: text("text").notNull(),
  done: boolean("done").default(false).notNull(),
});

export const sessions = pgTable("sessions", {
  id: text("id").primaryKey(),
  userId: text("user_id").notNull().unique(),
  displayName: text("display_name"),
  accessToken: text("access_token").notNull(),
  refreshToken: text("refresh_token"),
  expiresAt: timestamp("expires_at").notNull(),
  createdAt: timestamp("created_at").defaultNow().notNull(),
  lastLoginAt: timestamp("last_login_at").defaultNow().notNull(),
  loggedOutAt: timestamp("logged_out_at"),
});

export const assessmentProgress = pgTable("assessment_progress", {
  id: text("id").primaryKey(),
  userId: text("user_id").notNull(),
  exerciseId: text("exercise_id").notNull(),
  code: text("code").notNull(),
  passed: boolean("passed").notNull(),
  output: text("output").notNull(),
  attempts: integer("attempts").default(1).notNull(),
  updatedAt: timestamp("updated_at").defaultNow().notNull(),
});

export const quizProgress = pgTable("quiz_progress", {
  id: text("id").primaryKey(),
  userId: text("user_id").notNull(),
  categoryId: text("category_id").notNull(),
  passed: boolean("passed").notNull(),
  score: integer("score").notNull(),
  total: integer("total").notNull(),
  submittedAt: timestamp("submitted_at").defaultNow().notNull(),
});
