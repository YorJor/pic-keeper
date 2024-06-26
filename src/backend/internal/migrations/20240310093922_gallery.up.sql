CREATE TABLE galleries (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  photographer_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  name varchar(255) NOT NULL,
  avg_rating REAL,
  price integer NOT NULL,
  location varchar(255) NOT NULL,
  hours integer NOT NULL,
  description varchar(2000),
  delivery_time integer NOT NULL,
  included varchar(255) []
);


CREATE TYPE booking_status AS enum(
  'DRAFT',
  'USER_PAID',
  'CANCELLED',
  'C_REQ_CANCEL',
  'P_REQ_CANCEL',
  'COMPLETED',
  'PAID_OUT',
  'REQ_REFUND'
);


CREATE TABLE bookings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  customer_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  room_id UUID NOT NULL REFERENCES rooms (id) ON DELETE CASCADE,
  resulted_price integer NOT NULL,
  start_time timestamptz NOT NULL,
  end_time timestamptz NOT NULL,
  status booking_status NOT NULL DEFAULT 'USER_PAID',
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE photos (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  gallery_id UUID NOT NULL REFERENCES galleries (id) ON DELETE CASCADE,
  photo_key varchar(2000) NOT NULL
);


CREATE TABLE reviews (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  customer_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  booking_id UUID NOT NULL REFERENCES bookings (id) ON DELETE CASCADE,
  rating INTEGER NOT NULL,
  review_text varchar(2000)
);
