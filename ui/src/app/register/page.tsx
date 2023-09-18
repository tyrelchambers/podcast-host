"use client";
import { Button } from "@/components/ui/button";
import { Form, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { useForm } from "react-hook-form";
import { z } from "zod";

const registerSchema = z.object({
  email: z.string(),
  password: z.string(),
  confirmPassword: z.string(),
});

export default function Register() {
  const form = useForm({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const submitHandler = async (data: z.infer<typeof registerSchema>) => {
    await axios.post("http://localhost:8080/api/auth/register", data);
  };

  return (
    <div>
      <Form {...form}>
        <form
          className="flex flex-col gap-4"
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
              </FormItem>
            )}
          />

          <FormField
            name="password"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Password</Label>
                <Input type="password" placeholder="***" {...field} />
              </FormItem>
            )}
          />

          <FormField
            name="confirmPassword"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Confirm password</Label>
                <Input
                  type="password"
                  placeholder="Confirm your password"
                  {...field}
                />
              </FormItem>
            )}
          />

          <Button type="submit">Sign up</Button>
        </form>
      </Form>
    </div>
  );
}
