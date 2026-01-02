package o2odeadlock

import "o2o-deadlock/o2o-deadlock/pb"

func Foo() *pb.Outer {
	l := &pb.Outer{
		Either: &pb.Outer_Foo{
			Foo: "foo",
		},
	}

	return &pb.Outer{
		Either: l.GetEither(),
	}
}
