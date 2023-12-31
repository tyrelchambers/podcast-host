import { Button } from "@/components/ui/button";
import { Form, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import Link from "next/link";
import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const loginSchema = z.object({
  email: z.string(),
  password: z.string(),
});

const Login = () => {
  const form = useForm({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const submitHandler = async (data: z.infer<typeof loginSchema>) => {
    await axios.post("http://localhost:8080/api/auth/login", data, {
      headers: {
        "Content-Type": "application/json",
      },
      withCredentials: true,
    });
  };
  return (
    <div>
      <Form {...form}>
        <form
          className="flex flex-col gap-4 max-w-screen-sm mx-auto my-10"
          onSubmit={form.handleSubmit(submitHandler)}
        >
          <FormField
            name="email"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor="email">Email</Label>
                <Input
                  type="email"
                  placeholder="email@example.com"
                  {...field}
                />

                {form.formState.errors.email?.message && (
                  <span className="text-red-500">
                    {form.formState.errors.email?.message}
                  </span>
                )}
              </FormItem>
            )}
          />

          <FormField
            name="password"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Password</Label>
                <Input type="password" placeholder="***" {...field} />

                {form.formState.errors.password?.message && (
                  <span className="text-red-500">
                    {form.formState.errors.password?.message}
                  </span>
                )}
              </FormItem>
            )}
          />

          <Separator />

          <div>
            <p className="text-muted-foreground font-light text-sm">
              Haven&apos;t created an account yet?{" "}
              <Link href="/register" className="text-blue-500 underline">
                Sign up.
              </Link>
            </p>
          </div>
          <Button type="submit">Login</Button>
        </form>
      </Form>
    </div>
  );
};

export default Login;
