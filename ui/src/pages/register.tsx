"use client";
import { Button } from "@/components/ui/button";
import { Form, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import Link from "next/link";
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
    if (!form.getValues("confirmPassword")) {
      form.setError("confirmPassword", {
        message: "Please confirm your password",
      });
      return;
    }

    if (data.password !== data.confirmPassword) {
      form.setError("confirmPassword", {
        message: "Passwords do not match",
      });
      return;
    }

    await axios
      .post("http://localhost:8080/api/auth/register", data, {
        headers: {
          "Content-Type": "application/json",
        },
        withCredentials: true,
      })
      .then((res) => {
        console.log(res);
        if (res.status === 200) {
          // setCookie("token", res.headers.get("Authorization"));
        }
      });
  };

  return (
    <div>
      <Form {...form}>
        <form
          className="flex flex-col gap-4 max-w-screen-sm mx-auto my-10 bg-card p-4 rounded-xl"
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

                {form.formState.errors.confirmPassword?.message && (
                  <span className="text-red-500">
                    {form.formState.errors.confirmPassword?.message}
                  </span>
                )}
              </FormItem>
            )}
          />
          <Separator />

          <div>
            <p className="text-muted-foreground font-light text-sm">
              Already have an account?{" "}
              <Link href="/login" className="text-blue-500 underline">
                Login.
              </Link>
            </p>
          </div>
          <Button type="submit">Sign up</Button>
        </form>
      </Form>
    </div>
  );
}
